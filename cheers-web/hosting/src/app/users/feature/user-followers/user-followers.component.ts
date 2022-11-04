import { Component, OnInit } from '@angular/core';
import {User} from "../../../shared/data/models/user.model";

@Component({
  selector: 'app-user-followers',
  templateUrl: './user-followers.component.html',
  styleUrls: ['./user-followers.component.scss']
})
export class UserFollowersComponent implements OnInit {

    followers = [
        new User(),
        new User(),
        new User(),
        new User(),
    ]
  constructor() { }

  ngOnInit(): void {
  }

}
