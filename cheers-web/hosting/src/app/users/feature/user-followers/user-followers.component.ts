import { Component, OnInit } from '@angular/core';
import {User} from "../../../shared/data/models/user.model";
import {UserService} from "../../../shared/data/services/user.service";
import {firstValueFrom} from "rxjs";
import {UserItem} from "../../../../gen/ts/cheers/type/user/user";

@Component({
  selector: 'app-user-followers',
  templateUrl: './user-followers.component.html',
  styleUrls: ['./user-followers.component.scss']
})
export class UserFollowersComponent implements OnInit {

    followers: User[] = []
    following: User[] = []

  constructor(
      private userService: UserService,
  ) { }

  async ngOnInit()  {
        const user = await firstValueFrom(this.userService.user$)
     const followers = await firstValueFrom(this.userService.listFollowers(user.id))
      this.followers = followers
      const following = await firstValueFrom(this.userService.listFollowing(user.id))
      this.following = following
  }

}
