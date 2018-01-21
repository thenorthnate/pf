#Plotting Class for cotton application
#Nathan North

import tkinter

class Plot:
    def __init__(self):
        self.data = None

    def draw_something(self, canvas):
        canvas.create_arc(0, 0, 50, 50, fill="#FF8500", outline="#0776A0", style=tkinter.PIESLICE, start=240, extent=180)
