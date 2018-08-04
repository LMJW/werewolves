import pyglet
from cocos.director import director
from cocos.layer import Layer
from cocos.scene import Scene
from cocos.menu import Menu, MenuItem
from pyglet import gl


class Background(Layer):
    def __init__(self, resource_path=''):
        super().__init__()
        self.background = pyglet.resource.image('game_backgroup.png')

    def draw(self):
        self.background.blit(0, 0)


class Gamedisplay(Layer):
    def __init__(self):
        width, height = director.get_window_size()
        super().__init__()
        self.display = pyglet.graphics.Batch()
        background = pyglet.graphics.OrderedGroup(0)
        foreground = pyglet.graphics.OrderedGroup(1)
        char_image = pyglet.resource.image("villager_150x.png")
        chat_box = pyglet.resource.image("chat_box.png")
        self.char_sprite = pyglet.sprite.Sprite(
            char_image,
            x=40,
            y=height - char_image.height - 20,
            batch=self.display,
            group=background)
        self.chat_box = pyglet.sprite.Sprite(
            chat_box,
            x=60 + char_image.width,
            y=height - char_image.height - 20,
            batch=self.display,
            group=background)
        char_text = pyglet.text.Label(
            "Villager",
            font_name='Nemo Nightmares',
            font_size=40,
            color=(255, 50, 0, 255),
            x=120 + char_image.width,
            y=height - char_image.height // 2,
            batch=self.display,
            group=foreground)

    def draw(self):
        self.display.draw()

class ControlMenu(Menu):
    def __init__(self):
        super().__init__()
        m = []
        m.append(ControlMenuItem("Pause", self.pause_game))
        m.append(ControlMenuItem("Stop", self.stop_game))
        self.create_menu(m,layout_strategy=self.layout)

    def layout(self, menu):
        width, height = director.get_window_size()
        button_pos_x = (width//3, width*2//3)
        button_pos_y = 20
        font_item = dict(
            font_name='Nemo Nightmares',
            font_size=60,
            color=(255, 50, 0, 255),
            anchor_x='center',
            anchor_y='center'
            )
        font_item_selected = dict(
            font_name='Nemo Nightmares',
            font_size=60,
            color=(255, 180, 0, 255),
            anchor_x='center',
            anchor_y='center'
            )
        for i, c in self.children:
            font_item.update({'text':c.label})
            font_item_selected.update({'text':c.label})
            c.generateWidgets(button_pos_x[i], button_pos_y, font_item, font_item_selected)

    def pause_game(self):
        pass

    def stop_game(self):
        pass

class ControlMenuItem(MenuItem):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def generateWidgets(self, pos_x, pos_y, font_item, font_item_selected):
        self.button_batch = pyglet.graphics.Batch()
        self.button_batch_selected = pyglet.graphics.Batch()
        background = pyglet.graphics.OrderedGroup(0)
        foreground = pyglet.graphics.OrderedGroup(1)
        button_frame = pyglet.resource.image('control_button.png')
        button_frame_selected = pyglet.resource.image('control_button.png')
        self.button_frame = pyglet.sprite.Sprite(
            button_frame,
            x=pos_x - button_frame.width // 2,
            y=pos_y,
            batch=self.button_batch,
            group=background)
        self.button_frame_selected = pyglet.sprite.Sprite(
            button_frame_selected,
            x=pos_x - button_frame_selected.width // 2,
            y=pos_y,
            batch=self.button_batch_selected,
            group=background)
        self.item = pyglet.text.Label(
            **font_item,
            x=pos_x,
            y=pos_y+button_frame.height//2,
            batch=self.button_batch,
            group=foreground)
        self.item_selected = pyglet.text.Label(
            **font_item_selected,
            x=pos_x,
            y=pos_y+button_frame.height//2,
            batch=self.button_batch_selected,
            group=foreground)

    def draw(self):
        if self.is_selected:
            self.button_batch_selected.draw()
        else:
            self.button_batch.draw()

if __name__ == "__main__":
    import os
    root_dir = os.chdir("../")
    pyglet.resource.path = [os.getcwd() + '/data/image/']
    pyglet.resource.reindex()
    pyglet.font.add_directory(os.getcwd() + '/data/Fonts/')
    director.init(width=800, height=600)
    scene = Scene(Background(), Gamedisplay(), ControlMenu())
    director.run(scene)
