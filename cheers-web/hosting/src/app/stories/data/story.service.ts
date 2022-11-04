import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";

@Injectable({
    providedIn: 'root'
})
export class StoryService {

    constructor(
        private api: ApiService,
    ) {
    }

    getStoryFeed() {
        return this.api.getStoryFeed()
    }

    seenStory(storyId: string) {
        return this.api.seenStory(storyId).subscribe()
    }
}
