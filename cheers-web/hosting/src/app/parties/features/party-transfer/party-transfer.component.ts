import {Component, ElementRef, Inject, OnInit, ViewChild} from '@angular/core';
import {debounceTime, distinctUntilChanged, fromEvent, map, Observable, of} from "rxjs";
import {UserModel} from "../../../shared/data/models/user.model";
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../data/party.service";
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
    selector: 'app-party-transfer',
    templateUrl: './party-transfer.component.html',
    styleUrls: ['./party-transfer.component.sass']
})
export class PartyTransferComponent implements OnInit {
    isLoading: boolean = false
    selectedUser: UserModel | null;
    $searchResults: Observable<UserModel[] | null> = of(null)
    @ViewChild('searchInput', {static: true}) searchInput!: ElementRef;

    constructor(
        public dialogRef: MatDialogRef<PartyTransferComponent>,
        @Inject(MAT_DIALOG_DATA) public data: string,
        private userService: UserService,
        private partyService: PartyService,
        private snackBar: MatSnackBar,
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

    onUserClick(user: UserModel) {
        this.selectedUser = user;
    }

    onImgError($event: ErrorEvent) {

    }

    async onTransferClick() {
        const userId = this.selectedUser?.id
        if (!userId) {
            return
        }
        this.isLoading = true
        await this.partyService.transferParty(userId, this.data)
        this.snackBar.open("Party transferred successfully", "Hide");
        this.dialogRef.close()
        this.isLoading = false
    }
}
