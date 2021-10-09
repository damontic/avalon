class Card {
    constructor(frontFilepath, backFilepath, x, y){
        this.front = loadImage(frontFilepath);
        this.back = loadImage(backFilepath);
        this.resize = 0.25;
        this.imageToDraw = this.front;
        
        this.x = x;
        this.y = y;

        this.currentRotateSpeed = 0;
        this.rotateSpeed = 0.05;
        this.angle = 0;
        this.previousSinAngle = 1;
    }
    rotate() {
        this.currentRotateSpeed = this.rotateSpeed;
    }
    draw() {
        push();
        angleMode(RADIANS);

        if (this.currentRotateSpeed != 0) {
            this.angle += this.currentRotateSpeed;
            let sinAngle = sin(this.angle);
            if (sinAngle * this.previousSinAngle < 0) {
                this.currentRotateSpeed = 0;
            }
            this.previousSinAngle = sinAngle;
        }

        let cosAngle = cos(this.angle);
        if( cosAngle >= 0) {
            this.imageToDraw = this.front;
        } else {
            this.imageToDraw = this.back;
        }

        let adjustedWidth = this.imageToDraw.width * this.resize;
        let adjustHeight = this.imageToDraw.height * this.resize;

        imageMode(CENTER);
        rectMode(CENTER);
        translate(this.x, this.y);
        rotateY(this.angle);
        image(this.imageToDraw, 0, 0, adjustedWidth, adjustHeight);
        pop();
    }
    beingClicked(mouseX, mouseY) {
        let x0 = this.x - this.imageToDraw.width * this.resize / 2;
        let y0 = this.y - this.imageToDraw.height * this.resize / 2;
        let x1 = this.x + this.imageToDraw.width * this.resize / 2;
        let y1 = this.y + this.imageToDraw.height * this.resize / 2;

        if ( mouseX > x0 && mouseX < x1 && mouseY > y0 && mouseY < y1) {
            return true;
        }
        return true;
    }
}
