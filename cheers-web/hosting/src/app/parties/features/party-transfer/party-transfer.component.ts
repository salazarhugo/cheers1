import {Component, ElementRef, Inject, OnInit, ViewChild} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {debounceTime, distinctUntilChanged, fromEvent, map, Observable, of} from "rxjs";
import {User} from "../../../shared/data/models/user.model";
import {UserService} from "../../../shared/data/services/user.service";
import {UserItem} from "../../../../gen/ts/cheers/type/user/user";

@Component({
    selector: 'app-party-transfer',
    templateUrl: './party-transfer.component.html',
    styleUrls: ['./party-transfer.component.sass']
})
export class PartyTransferComponent implements OnInit {

    selectedUser: UserItem | null;
    $searchResults: Observable<UserItem[] | null> = of(null)
    @ViewChild('searchInput', {static: true}) searchInput!: ElementRef;

    constructor(
        public dialogRef: MatDialogRef<PartyTransferComponent>,
        @Inject(MAT_DIALOG_DATA) public data: any,
        private userService: UserService,
    ) { }

    ngOnInit(): void {
        fromEvent(this.searchInput.nativeElement, 'keyup').pipe(
            map((event: any) => event.target.value),
            debounceTime(500),
            distinctUntilChanged()
        ).subscribe((query: string) => {
            this.searchUser(query)
        });
    }

    searchUser(query: string) {
        this.$searchResults = this.userService.searchUser(query)
    }

    onNoClick() {
        this.dialogRef.close();
    }

    onUserClick(user: UserItem) {
        this.selectedUser = user;
    }

    onImgError($event: ErrorEvent) {

    }
}
