import pyglet
from cocos.director import director
from cocos.layer import ColorLayer, Layer
from cocos.menu import Menu, MenuItem, MultipleMenuItem
from cocos.scene import Scene
from cocos.text import Label
from pyglet.window import key, mouse

from config import (_DEFAULT_SELECTED_TEXT_COLOR, _DEFAULT_TEXT_COLOR,
                    _GAME_CHARACTERS)


def menuLayout(menu):
    width, height = director.get_window_size()
    fo = pyglet.font.load(menu.font_item['font_name'],
                          menu.font_item['font_size'])
    fo_height = int((fo.ascent - fo.descent) * 0.5)
    if menu.menu_halign == pyglet.font.Text.CENTER:
        pos_x = width // 2
    elif menu.menu_halign == pyglet.font.Text.RIGHT:
        pos_x = width - menu.menu_hmargin
    elif menu.menu_halign == pyglet.font.Text.LEFT:
        pos_x = menu.menu_hmargin
    else:
        raise Exception("Invalid anchor_x value for menu")

    for idx, i in enumerate(menu.children):
        item = i[1]
        if menu.menu_valign == pyglet.font.Text.CENTER:
            pos_y = (height + (len(menu.children) - 2 * idx) * fo_height -
                     menu.title_height) * 0.5
        elif menu.menu_valign == pyglet.font.Text.TOP:
            pos_y = (height - ((idx + 0.8) * fo_height) - menu.title_height -
                     menu.menu_vmargin)
        elif menu.menu_valign == pyglet.font.Text.BOTTOM:
            pos_y = (
                0 + fo_height * (len(menu.children) - idx) + menu.menu_vmargin)
        item.transform_anchor = (pos_x, pos_y)
        item.generateWidgets(pos_x, pos_y, menu.font_html, menu.font_html)


class Setting(Layer):
    def __init__(self):
        super().__init__()


class CharacterMenuItem(MultipleMenuItem):
    def __init__(self, *args, **kwargs):
        self.text_template = "<pre><font color={2}>{0:<20}:{1:>3}</font></pre>"
        super().__init__(*args, **kwargs)

    # overwrite the generateWidgets to get the fixed width label
    def generateWidgets(self, pos_x, pos_y, font_item, font_item_selected):
        font_item['x'] = int(pos_x)
        font_item['y'] = int(pos_y)
        font_item['text'] = self.label

        self.item = pyglet.text.HTMLLabel(**font_item)

        font_item_selected['x'] = int(pos_x)
        font_item_selected['y'] = int(pos_y)
        font_item_selected['text'] = self._get_selected_label()

        self.item_selected = pyglet.text.HTMLLabel(**font_item_selected)

    def _get_label(self):
        return self.text_template.format(self.my_label, self.idx,
                                         _DEFAULT_TEXT_COLOR)

    def _get_selected_label(self):
        return self.text_template.format(self.my_label, self.idx,
                                         _DEFAULT_SELECTED_TEXT_COLOR)

    def on_key_press(self, symbol, modifiers):
        if symbol == key.LEFT:
            self.idx = max(0, self.idx - 1)
        elif symbol in (key.RIGHT, key.ENTER):
            self.idx = min(len(self.items) - 1, self.idx + 1)

        if symbol in (key.LEFT, key.RIGHT, key.ENTER):
            self.item.text = self._get_label()
            self.item_selected.text = self._get_selected_label()
            self.callback_func(self.idx)
            return {self.my_label:self.idx}


class CharactersMenu(Menu):
    is_event_handler = True

    def __init__(self):
        self.title_template = "Total game characters : {0:>3}"
        self._game_char_dict = dict()
        super().__init__()

        self.font_title.update({'font_size': 30})
        self.font_html = {
            'anchor_y': 'center',
            'anchor_x': 'center',
            'dpi': 96,
        }

        m = []
        for k, v in _GAME_CHARACTERS.items():
            m.append(CharacterMenuItem(k, self.haha, v))

        self.create_menu(m, layout_strategy=menuLayout)

    def _activate_item(self):
        if self.activate_sound:
            self.activate_sound.play()
        self.children[self.selected_index][1].on_activated()
        updated_player_data = self.children[self.selected_index][1].on_key_press(key.ENTER, 0)
        if updated_player_data:
            self._game_char_dict.update(updated_player_data)
            self._generate_title()
            return True

    def _generate_title(self):
        width, height = director.get_window_size()

        self.font_title['x'] = width // 2
        self.font_title['text'] = self.title_template.format(sum(self._game_char_dict.values()))
        self.title_label = pyglet.text.Label(**self.font_title)
        self.title_label.y = height - self.title_label.content_height // 2

        self.title_height = self.title_label.content_height

    def on_key_press(self, symbol, modifiers):
        if symbol == key.ESCAPE:
            self.on_exit()
            return True
        elif symbol in (key.ENTER, key.NUM_ENTER):
            self._activate_item()
            return True
        elif symbol in (key.DOWN, key.UP):
            if symbol == key.DOWN:
                new_idx = self.selected_index + 1
            elif symbol == key.UP:
                new_idx = self.selected_index - 1

            if new_idx < 0:
                new_idx = len(self.children) - 1
            elif new_idx > len(self.children) - 1:
                new_idx = 0
            self._select_item(new_idx)
            return True
        else:
            # send the menu item the rest of the keys
            updated_player_data = self.children[self.selected_index][1].on_key_press(symbol, modifiers)
            try:
                self._game_char_dict.update(updated_player_data)
            except TypeError:
                return
            # play sound if key was handled
            if updated_player_data and self.activate_sound:
                self.activate_sound.play()
            self._generate_title()
            return True

    def haha(self, *args):
        print(args)



if __name__ == '__main__':
    director.init(width=728, height=424)
    scene = Scene(Setting(), CharactersMenu())
    director.run(scene)
