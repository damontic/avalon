class AvalonText {
    constructor(imageFile, x, y) {
        this.x = x;
        this.y = y;
        this.imageText = loadImage(imageFile);
    }
    draw() {
        push();
        imageMode(CORNER);
        stroke(0);
        strokeWeight(5);
        fill(255,255,255,255);
        translate(this.x, this.y);
        rect(0,0,this.imageText.width, this.imageText.height);
        image(this.imageText, 0, 0);
        pop();
    }
}
