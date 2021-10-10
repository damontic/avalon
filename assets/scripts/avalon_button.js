class AvalonButton {
    constructor(avalon, text, x, y, font, font_size = 24) {
        this.avalon = avalon
        this.text = text
        this.left_margin = 5
        this.right_margin = 5
        this.upper_margin = 5
        this.lower_margin = 5
        this.x = x + this.left_margin
        this.y = y
        this.font_size = font_size
        this.font = font
        this.bbox = font.textBounds(this.text, this.x, this.y, this.font_size)
        this.bbox.x = this.bbox.x - this.left_margin
        this.bbox.y = this.bbox.y + this.upper_margin
        this.bbox.w = this.bbox.w + this.right_margin + this.left_margin
        this.bbox.h = this.bbox.h + this.lower_margin + this.upper_margin
    }
    draw() {
        push()
        rect(this.bbox.x, this.bbox.y, this.bbox.w, this.bbox.h)
        fill(WHITE);
        stroke(0);
        noStroke()
        fill(BLUE)
        noStroke()
        textFont(this.font)
        textSize(this.font_size)
        text(this.text, this.x, this.y)
        pop()
    }
}
