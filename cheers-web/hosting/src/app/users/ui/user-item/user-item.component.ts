import {Component, Input, OnInit} from '@angular/core';
import {User} from "../../../shared/data/models/user.model";

@Component({
    selector: 'app-user-item',
    templateUrl: './user-item.component.html',
    styleUrls: ['./user-item.component.sass']
})
export class UserItemComponent implements OnInit {

    @Input() user: User

    constructor() {
    }

    ngOnInit(): void {
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }
}
