import pyglet

window = pyglet.window.Window()



car_image = pyglet.resource.image('car.png')

car_image_2 = pyglet.resource.image('car2.png')


batch = pyglet.graphics.Batch()
car = pyglet.sprite.Sprite(car_image, x=100, y =100, batch=batch)
# car_2 = pyglet.sprite.Sprite(car_image_2, x=100, y=100, batch=batch)

label = pyglet.text.Label('Woops',
                          font_name='Times New Roman',
                          font_size=36,
                          x=window.width//2, y=window.height//2,
                          anchor_x='center', anchor_y='center', batch=batch)

@window.event
def on_draw():
    window.clear()
    label.draw()
    batch.draw()

pyglet.app.run()