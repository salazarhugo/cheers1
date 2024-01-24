import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {UserModel} from "../../../shared/data/models/user.model";

@Component({
    selector: 'app-user-item',
    templateUrl: './user-item.component.html',
    styleUrls: ['./user-item.component.sass']
})
export class UserItemComponent implements OnInit {

    @Input() user: UserModel
    @Input() isSelected: boolean = false;
    @Output() onClick = new EventEmitter<UserModel>()

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
