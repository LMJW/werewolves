import pyglet
from cocos.layer import Layer
from cocos.scene import Scene
from cocos.director import director
class Background(Layer):
    def __init__(self):
        super().__init__()
        self.image = pyglet.resource.image('background.jpg')

    def draw(self):
        self.image.blit(0,0)

class Gamedisplay(Layer):
    def __init__(self):
        super().__init__()
        self.display = pyglet.graphics.Batch()

if __name__ == "__main__":
    director.init(width=728, height=424)
    scene = Scene(Background())
    director.run(scene)
