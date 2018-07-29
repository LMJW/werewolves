import pyglet
from cocos.director import director
from cocos.layer import Layer
from cocos.scene import Scene
from cocos.menu import Menu, MenuItem


class Background(Layer):
    def __init__(self):
        super().__init__()
        self.image = pyglet.resource.image('background.jpg')

    def draw(self):
        self.image.blit(0,0)

class MenuButton(MenuItem):
    is_event_handler = True
    def __init__(self, label, callback_func):
        super().__init__(label, callback_func)

    def on_key_press(self, symbol, modifiers):
        if symbol & pyglet.window.mouse.LEFT:
            print(self.label +" is clicked")
            print(self.callback_func())

class StartUpMenu(Menu):
    is_event_handler = True
    def __init__(self):
        super().__init__()
        m = []
        m.append(MenuButton("Setting", self.on_game_setting))
        m.append(MenuButton("Start", self.on_new_game))
        self.create_menu(m)

    def on_new_game(self):
        print("start game clicked")

    def on_game_setting(self):
        print("Game setting clicked")

if __name__ == '__main__':
    director.init(width=728, height=424)
    scene = Scene(Background(), StartUpMenu())
    director.run(scene)