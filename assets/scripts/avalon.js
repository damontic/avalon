class Avalon {
    constructor(width, height, params){
        this.room = params.room ? params.room : null
        this.width = width
        this.height = height
        this.left_margin = 10
        this.right_margin = 10
        this.upper_margin = 10
        this.lower_margin = 10
        this.state = "menu"
        // backgrounds from farthest to closer
        this.background_zs = [-500]

        this.title = new AvalonImage(this, IMAGE_TITLE, this.left_margin, this.upper_margin)

        this.menu = new AvalonMenu(
            this,
            this.left_margin,
            this.title.height + this.upper_margin,
            this.width - this.left_margin - this.right_margin,
            this.height - this.title.height - this.upper_margin - this.lower_margin
        )

        this.room_waiting = new AvalonRoomWaiting(this)
        this.avalon_team_discussion = new AvalonTeamDiscussion(this)
        this.avalon_quest = new AvalonQuest(this)

        this.merlin_card = new AvalonCard(IMAGE_MERLIN, IMAGE_CHARACTER_CARD_BACK, 0, 0)
        this.morgana_card = new AvalonCard(IMAGE_MORGANA, IMAGE_CHARACTER_CARD_BACK, 150, 0)
    }
    draw() {
        background(BLACK)
        push()
        // Sets the origin to the upper left of the canvas
        // Add some margin space
        translate(
            - this.width / 2,
            - this.height / 2
        )

        this.title.draw()

        switch(this.state) {
            case "menu":
                this.menu.draw()
                break
            case "room_waiting":
                this.room_waiting.draw()
                break
            case "avalon_team_discussion":
                this.avalon_team_discussion.draw()
                break
            case "avalon_quest":
                this.avalon_quest.draw()
                break
            default:
                console.error("[ERROR]: Undefined state in avalon.js")
                this.state = "menu"
                break
        }

        this.merlin_card.draw()
        this.morgana_card.draw()
        pop()
    }
}
