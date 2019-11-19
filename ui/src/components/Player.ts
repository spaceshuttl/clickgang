export default class Player {
    id: string
    initialised: boolean

    constructor(id?: string) {
        if (id) {
            this.id = id
            this.initialised = true
        } else {
            this.initialised = false
        }
    }
}