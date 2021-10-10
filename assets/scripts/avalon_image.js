class AvalonImage {
    constructor(avalon, image, x, y, resizeConfig = null) {
        this.avalon = avalon
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
        translate(0,0,this.avalon.background_zs[0])
        if(this.resizeConfig != null)
            this.image.resize(this.resizeConfig.width, this.resizeConfig.height)
        image(this.image, this.x, this.y)
        pop()
    }
}
