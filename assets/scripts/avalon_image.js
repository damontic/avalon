class AvalonImage {
    constructor(image, x, y, resizeConfig = null) {
        this.x = x
        this.y = y
        this.resizeConfig = resizeConfig
        this.imageMode = imageMode
        this.image = image
        this.height = image.height
        this.width = image.width
    }
    draw() {
        push()
        if(this.resizeConfig != null)
            this.image.resize(this.resizeConfig.width, this.resizeConfig.height)
        image(this.image, this.x, this.y)
        pop()
    }
}
