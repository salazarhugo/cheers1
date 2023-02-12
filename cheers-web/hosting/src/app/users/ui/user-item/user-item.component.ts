import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {User} from "../../../shared/data/models/user.model";
import {UserItem} from "../../../../gen/ts/cheers/type/user/user";

@Component({
    selector: 'app-user-item',
    templateUrl: './user-item.component.html',
    styleUrls: ['./user-item.component.sass']
})
export class UserItemComponent implements OnInit {

    @Input() user: UserItem
    @Input() isSelected: boolean = false;
    @Output() onClick = new EventEmitter<UserItem>()

    constructor() {
    }

    click() {
        this.onClick.emit(this.user)
    }

    ngOnInit(): void {
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }
}
