import {Component, Input, OnInit} from '@angular/core';
import {User} from "../../../shared/data/models/user.model";

@Component({
  selector: 'app-profile-header',
  templateUrl: './profile-header.component.html',
  styleUrls: ['./profile-header.component.sass']
})
export class ProfileHeaderComponent implements OnInit {

  @Input() user: User

  constructor() {
  }

  ngOnInit(): void {
  }

  onImgError(event: any) {
    event.target.src = 'assets/default_profile_picture.png';
  }
}
