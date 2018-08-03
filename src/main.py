import cocos

class Helloworld(cocos.layer.Layer):
    def __init__(self):
        super().__init__()
        label = cocos.text.Label(
            "hello, world",
            font_size = 32,
            anchor_x = 'center',
            anchor_y = 'center'
        )
        label.position = 320, 240
        self.add(label)

cocos.director.director.init()
hello_layer = Helloworld()
main_scene = cocos.scene.Scene (hello_layer)
cocos.director.director.run(main_scene)
