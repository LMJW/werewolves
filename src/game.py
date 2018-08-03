import pyglet
from cocos.director import director
from cocos.layer import Layer
from cocos.scene import Scene
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


if __name__ == "__main__":
    import os
    root_dir = os.chdir("../")
    pyglet.resource.path = [os.getcwd() + '/data/image/']
    pyglet.resource.reindex()
    pyglet.font.add_directory(os.getcwd() + '/data/Fonts/')
    director.init(width=800, height=600)
    scene = Scene(Background(), Gamedisplay())
    director.run(scene)
