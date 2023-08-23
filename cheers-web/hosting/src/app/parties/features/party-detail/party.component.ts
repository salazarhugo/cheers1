import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {Party} from "../../../shared/data/models/party.model";
import {PartyService} from "../../data/party.service";
import {Clipboard} from '@angular/cdk/clipboard';
import {MatSnackBar} from "@angular/material/snack-bar";
import {UserService} from "../../../shared/data/services/user.service";
import {firstValueFrom} from "rxjs";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {MatDialog} from "@angular/material/dialog";
import {PartyInviteComponent} from "../../ui/party-invite/party-invite.component";
import {WatchStatus} from "../../../../gen/ts/cheers/party/v1/party_service";
import {PostDeleteDialogComponent} from "../../../posts/ui/post-delete-dialog/post-delete-dialog.component";
import {PartyTransferComponent} from "../party-transfer/party-transfer.component";
import {PartyDuplicateComponent} from "../party-duplicate/party-duplicate.component";
import {User} from "../../../shared/data/models/user.model";

@Component({
    selector: 'app-party-detail',
    templateUrl: './party.component.html',
    styleUrls: ['./party.component.sass']
})
export class PartyComponent implements OnInit {

    party: Party | undefined
    partyId: string | null
    viewer: User | null

    constructor(
        public router: Router,
        private route: ActivatedRoute,
        private partyService: PartyService,
        private clipboard: Clipboard,
        private snackBar: MatSnackBar,
        public userService: UserService,
        private afAuth: AngularFireAuth,
        private matDialog: MatDialog,
    ) {
    }

    ngOnInit(): void {
        // From PartyResolver
        this.party = this.route.snapshot.data['party'] as Party;
    }

    onInterestedClick() {
        if (this.party!.interested)
            this.partyService.uninterested(this.partyId!).subscribe(res => console.log(res))
        else
            this.partyService.interested(this.partyId!).subscribe(res => console.log(res))

        this.party!.interested = !this.party!.interested
    }

    copyLink() {
        this.clipboard.copy(window.location.href);
        this.snackBar.open("Link copied to clipboard", "Hide", {
            duration: 3000
        })
    }

    onInviteClick() {
        this.matDialog.open(PartyInviteComponent)
    }

    async onGoingToggle() {
        const user = await firstValueFrom(this.afAuth.authState)
        if (user == null) {
            await this.router.navigate(['sign-in'])
        }
        const party = this.party!!
        if (party.going) {
            party.going = false
            await this.partyService.answerParty(this.partyId!, WatchStatus.UNWATCHED)
        }
        else {
            party.going = true
            await this.partyService.answerParty(this.partyId!, WatchStatus.GOING)
        }
    }

    onDuplicateClick() {
        const dialogRef = this.matDialog.open(PartyDuplicateComponent, {
            panelClass: 'cheers-dialog'
        });

        dialogRef.afterClosed().subscribe(async result => {
            if (result == false)
                return;

            await firstValueFrom(this.partyService.duplicateParty(this.partyId!));
            this.snackBar.open("Duplicated party", "Hide");
        });
    }

    onTransferClick() {
        const dialogRef = this.matDialog.open(PartyTransferComponent, {
            panelClass: 'cheers-dialog'
        });

        dialogRef.afterClosed().subscribe(async userId => {
            if (userId == "")
                return;

            await this.partyService.transferParty(userId, this.partyId!)
            this.snackBar.open("Transfered party", "Hide");
        });
    }
}
