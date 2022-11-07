import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap} from "@angular/router";
import {StoryService} from "../../data/story.service";
import {Observable, of} from "rxjs";
import {Story} from "../../../shared/data/models/story.model";
import {Location} from '@angular/common';


@Component({
    selector: 'app-story-feed',
    templateUrl: './stories.component.html',
    styleUrls: ['./stories.component.sass']
})
export class StoriesComponent implements OnInit {

    selectedStory: Story | null = null
    story$: Observable<Story[] | null> = of(null)
    isLoading: boolean = true

    constructor(
        private storyService: StoryService,
        private route: ActivatedRoute,
        private location: Location,
    ) {
    }

    ngOnInit(): void {
        this.route.paramMap.subscribe((params: ParamMap) => {
            const storyId = params.get("id")
            this.story$ = this.storyService.getStoryFeed()
            this.story$.subscribe(stories => {
                const story = stories?.find(s => s.id == storyId)
                if (story) {
                    this.selectedStory = story
                    this.storyService.seenStory(story.id)
                }
            })
        })
    }

    back() {
        this.location.back()
    }

    onLoad() {
        this.isLoading = false;
    }
}
