import pyglet
from cocos.director import director
from cocos.layer import Layer, ColorLayer
from cocos.scene import Scene
from cocos.menu import Menu, MultipleMenuItem
from cocos.text import Label


def menuLayout(menu):
    width, height = director.get_window_size()
    fo = pyglet.font.load(menu.font_item['font_name'], menu.font_item['font_size'])
    fo_height = int((fo.ascent - fo.descent) * 0.9)

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
            pos_y = (height + (len(menu.children) - 2 * idx)
                     * fo_height - menu.title_height) * 0.5
        elif menu.menu_valign == pyglet.font.Text.TOP:
            pos_y = (height - ((idx + 0.8) * fo_height)
                     - menu.title_height - menu.menu_vmargin)
        elif menu.menu_valign == pyglet.font.Text.BOTTOM:
            pos_y = (0 + fo_height * (len(menu.children) - idx) +
                     menu.menu_vmargin)
        item.transform_anchor = (pos_x, pos_y)
        item.generateWidgets(pos_x, pos_y, menu.font_item,
                             menu.font_item_selected)

class Setting(Layer):
    def __init__(self):
        super().__init__()

class CharacterMenuItem(MultipleMenuItem):

    def __init__(self, *args, **kwargs):
        winx, _ = director.get_window_size()
        self.labelwidth = winx//20
        super().__init__(*args, **kwargs)

    # overwrite the generateWidgets to get the fixed width label
    def generateWidgets(self, pos_x, pos_y, font_item, font_item_selected):
        font_item['x'] = int(pos_x)
        font_item['y'] = int(pos_y)
        font_item['text'] = self.label
        # font_item['width'] = self.labelwidth
        # font_item['multiline'] = True
        self.item = pyglet.text.Label(**font_item)
        print(self.item.content_width)
        # self.item.content_width = self.labelwidth
        # print(self.item.content_width)

        font_item_selected['x'] = int(pos_x)
        font_item_selected['y'] = int(pos_y)
        font_item_selected['text'] = self.label
        # font_item_selected['width'] = self.labelwidth
        # font_item_selected['multiline'] = True

        self.item_selected = pyglet.text.Label(**font_item_selected)
        print(self.item_selected.content_width)
        self.item_selected.content_width = self.labelwidth
        print(self.item_selected.content_width)

    def _get_label(self):
        label_len = len(self.my_label)
        padding_space = ' '*(self.labelwidth-label_len)
        return self.my_label + padding_space + self.items[self.idx]

class CharactersMenu(Menu):
    is_event_handler = True
    def __init__(self):
        super().__init__()
        self.font_item = {
            'font_name': 'Arial',
            'font_size': 32,
            'bold': False,
            'italic': False,
            'anchor_y': 'center',
            'anchor_x': 'center',
            'color': (192, 192, 192, 255),
            'dpi': 96,
        }

        self.font_item_selected = {
            'font_name': 'Arial',
            'font_size': 32,
            'bold': True,
            'italic': False,
            'anchor_y': 'center',
            'anchor_x': 'center',
            'color': (255, 255, 255, 255),
            'dpi': 96,
        }

        m = []
        m.append(CharacterMenuItem("Setting", self.haha, [str(i) for i in range(10)]))
        m.append(CharacterMenuItem("Start", self.haha, ['0', '1', '2']))
        self.create_menu(m)

    def haha(self, *args):
        print('haha')


if __name__ == '__main__':
    director.init(width=728, height=424)
    scene = Scene(Setting(), CharactersMenu())
    director.run(scene)