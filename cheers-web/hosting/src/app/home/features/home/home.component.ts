import {Component, OnDestroy, OnInit} from '@angular/core';
import {UserService} from "../../../shared/data/services/user.service";
import {Observable, of, Subscription} from "rxjs";
import {Post} from "../../../shared/data/models/post.model";
import {Story} from "../../../shared/data/models/story.model";
import {StoryService} from "../../../stories/data/story.service";
import {User} from "../../../shared/data/models/user.model";
import {UntypedFormControl, Validators} from "@angular/forms";
import {PostService} from "../../../posts/data/post.service";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {PostCreateDialogComponent} from "../../../posts/ui/post-create-dialog/post-create-dialog.component";
import {MatDialog} from "@angular/material/dialog";

@Component({
    selector: 'app-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.sass']
})
export class HomeComponent implements OnInit, OnDestroy {

    isPosting = false
    private subs: Subscription[] = [];

    user: User | null = null
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
        public dialog: MatDialog,
    ) {
        this.subs.push(this.userService.user$.subscribe( user => this.user = user))
    }


    ngOnInit(): void {
        this.subs.push(this.postService.getPostFeed().subscribe(posts => {
            this.posts = posts
        }))
        this.story$ = this.storyService.getStoryFeed()
    }

    ngOnDestroy() {
        this.subs.forEach(sub => sub.unsubscribe())
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

        const index = this.posts.indexOf(postResponse, 0)
        if (index > -1) {
            this.posts.splice(index, 1)
        }
    }

    openCreatePostDialog() {
        const dialogRef = this.dialog.open(PostCreateDialogComponent, {
            panelClass: 'cheers-dialog'
        })

        dialogRef.afterClosed().subscribe(result => {
        })
    }
}
