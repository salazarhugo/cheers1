import {Component, Inject, OnInit} from '@angular/core';
import {User} from "../../../shared/data/models/user.model";
import {UserItem} from "../../../../gen/ts/cheers/type/user/user";
import {UserService} from "../../../shared/data/services/user.service";
import {firstValueFrom} from "rxjs";
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";

@Component({
    selector: 'app-party-invite',
    templateUrl: './party-invite.component.html',
    styleUrls: ['./party-invite.component.sass']
})
export class PartyInviteComponent implements OnInit {

    isLoading = false
    selectedUsers: Set<UserItem> = new Set()
    users: UserItem[]

    constructor(
        public dialogRef: MatDialogRef<PartyInviteComponent>,
        @Inject(MAT_DIALOG_DATA) public data: any,
        private userService: UserService,
    ) {
    }

    async ngOnInit()  {
        const user = await firstValueFrom(this.userService.user$)
        const users = await firstValueFrom(this.userService.listFollowers(user.id))
        this.users = users
    }

    onSelectUser($event: UserItem) {
        if (this.selectedUsers.has($event))
            this.selectedUsers.delete($event)
        else
            this.selectedUsers.add($event)
    }

    onSendInvitations() {
        this.isLoading = true
    }
}
