import {Component, OnInit} from '@angular/core';
import {AuthService} from "../../../shared/data/services/auth.service";
import {UserService} from "../../../shared/data/services/user.service";
import {Observable, of} from "rxjs";
import {UserModel} from "../../../shared/data/models/user.model";
import {Party} from "../../../shared/data/models/party.model";
import {PartyService} from "../../../parties/data/party.service";
import {Router} from "@angular/router";
import {ActivityService} from "../../../activities/data/activity.service";
import {PostService} from "../../../posts/data/post.service";
import {PostModel} from "../../../shared/data/models/post.model";

@Component({
    selector: 'app-profile',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.sass']
})
export class ProfileComponent implements OnInit {

    $user: Observable<UserModel | null> = of(null)
    $parties: Observable<Party[] | null> = of(null)
    $posts: Observable<PostModel[] | null> = of(null)

    constructor(
        private authService: AuthService,
        private userService: UserService,
        private eventService: PartyService,
        private postService: PostService,
        private activityService: ActivityService,
        private router: Router,
    ) {
        this.$user = this.userService.user$
        this.$parties = this.eventService.getPartyFeed()
        this.$posts = this.postService.getPostFeed()
    }

    ngOnInit(): void {
        this.activityService.listActivity()
    }

    async onEdit() {
        await this.router.navigate(['accounts/user-edit'])
    }

    async onSignOut() {
        await this.authService.signOut()
        await this.router.navigate(['sign-in'])
    }

    onPartyClick(partyId: string) {
    }
}
