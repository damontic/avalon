class AvalonTeamDiscussion {
    constructor() {
        let menu_background = floor(random(14))
        this.image_menu_background = new AvalonImage(`images/backgrounds/background${menu_background}.jpg`, CENTER, 0, 0)
        this.title = new AvalonImage(`images/text/title.png`, CENTER, 0, 0)
        this.image_players = new AvalonImage(`images/text/players.png`, CENTER, 0, 0)
        this.image_numbers = new AvalonImage(`images/numbers.png`, CENTER, 0, 0)
        this.one = {
            ox: 0,
            oy: 0,
            width: 203,
            height: 211
        }
    }
    draw() {
        this.image_menu_background.draw()
        //this.title.draw()

    }
}
