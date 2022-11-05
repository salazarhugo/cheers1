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

    isPosting = false
    private sub: any;

    $user: Observable<User | null> = of(null)
    posts: PostResponse[] | null = null
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
        this.sub = this.postService.getPostFeed().subscribe(posts => {
            this.posts = posts
        })
        this.story$ = this.storyService.getStoryFeed()
    }

    ngOnDestroy() {
        this.sub.unsubscribe();
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    createPost() {
        this.isPosting = true
        const post = new Post()
        post.caption = this.tweet.value
        this.postService.createPost(post).subscribe(post => {
            this.posts?.unshift(post)
            this.tweet.reset()
            this.isPosting = false
        })
    }

    onDelete(postResponse: PostResponse) {
        if (this.posts == undefined) return

        const index = this.posts.indexOf(postResponse, 0);
        if (index > -1) {
            this.posts.splice(index, 1);
        }
    }
}
