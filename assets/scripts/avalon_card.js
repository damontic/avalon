class AvalonCard {
    constructor(frontImage, backImage, x, y){
        this.front = frontImage
        this.back = backImage
        this.x = x
        this.y = y

        this.imageToDraw = this.front

        this.currentRotateSpeed = 0
        this.rotateSpeed = 0.05
        this.angle = 0
        this.previousSinAngle = 1
        this.resizeWidth = 150
    }
    rotate() {
        this.currentRotateSpeed = this.rotateSpeed
    }
    updateState() {
        if (this.currentRotateSpeed != 0) {
            this.angle += this.currentRotateSpeed
            let sinAngle = sin(this.angle)
            if (sinAngle * this.previousSinAngle < 0) {
                this.currentRotateSpeed = 0
            }
            this.previousSinAngle = sinAngle
        }

        let cosAngle = cos(this.angle)
        if( cosAngle >= 0) {
            this.imageToDraw = this.front
        } else {
            this.imageToDraw = this.back
        }
    }
    draw() {
        this.updateState()
        push();
        this.imageToDraw.resize(this.resizeWidth, 0)
        imageMode(CENTER)
        let proportion = this.imageToDraw.width / this.resizeWidth
        translate(this.x + this.resizeWidth / 2, this.y + this.imageToDraw.height / proportion / 2)
        rotateY(this.angle)
        image(this.imageToDraw, 0, 0)
        pop()
    }
    checkBeingClicked() {
        console.log("Checking Merlin clicked")
        let x0 = this.x - this.imageToDraw.width
        let y0 = this.y - this.imageToDraw.height
        let x1 = this.x + this.imageToDraw.width
        let y1 = this.y + this.imageToDraw.height

        if ( mouseX > x0 && mouseX < x1 && mouseY > y0 && mouseY < y1) {
            this.rotate()
        }
    }
}
