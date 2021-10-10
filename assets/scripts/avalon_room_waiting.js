class AvalonRoomWaiting {
    constructor() {
        let menu_background = floor(random(3))
        this.image_menu_background = new AvalonImage(`images/backgrounds/round_table/rt${menu_background}.jpg`, CENTER, 0, 0)
    }
    draw() {
        this.image_menu_background.draw()
    }
}
