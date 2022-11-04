import {Component, OnInit} from '@angular/core';
import {UserService} from "../../../shared/data/services/user.service";
import {Observable, of} from "rxjs";
import {Post} from "../../../shared/data/models/post.model";
import {Story} from "../../../shared/data/models/story.model";
import {StoryService} from "../../../stories/data/story.service";
import {User} from "../../../shared/data/models/user.model";
import {UntypedFormControl, Validators} from "@angular/forms";
import {PostService} from "../../../posts/data/post.service";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";

@Component({
    selector: 'app-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.sass']
})
export class HomeComponent implements OnInit {

    $user: Observable<User | null> = of(null)
    $posts: Observable<PostResponse[] | null> = of(null)
    story$: Observable<Story[] | null> = of(null)

    tweet = new UntypedFormControl('', Validators.compose([
            Validators.minLength(1),
            Validators.maxLength(300),
            Validators.required,
        ]),
    )

    constructor(
        private userService: UserService,
        private storyService: StoryService,
        private postService: PostService,
    ) {
        this.$user = this.userService.user$
    }

    ngOnInit(): void {
        this.$posts = this.userService.getPostFeed()
        this.story$ = this.storyService.getStoryFeed()
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    createPost() {
        const post = new Post()
        post.caption = this.tweet.value
        this.postService.createPost(post).subscribe(post => {
            console.log(post)
        })
    }
}
