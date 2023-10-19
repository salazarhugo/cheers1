import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {Party, WatchStatus} from "../../../shared/data/models/party.model";
import {PartyService} from "../../data/party.service";
import {Clipboard} from '@angular/cdk/clipboard';
import {MatSnackBar} from "@angular/material/snack-bar";
import {UserService} from "../../../shared/data/services/user.service";
import {firstValueFrom} from "rxjs";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {MatDialog} from "@angular/material/dialog";
import {PartyInviteComponent} from "../../ui/party-invite/party-invite.component";
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
    viewer: User | null
    watchStatusNames: WatchStatus[] =
        Object.keys(WatchStatus)
            .filter((v) => isNaN(Number(v)))
            .map((name) => {
                return WatchStatus[name as keyof typeof WatchStatus]
            });

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
        this.userService.user$.subscribe(res => this.viewer = res)
    }

    async onInterestedClick() {
        const partyId = this.party?.id
        if (partyId == undefined)
            return

        // this.party = undefined
        await this.partyService.answerParty(partyId, WatchStatus.INTERESTED)
        this.party!.watchStatus = WatchStatus.INTERESTED
        this.party!.isNotInterested = false
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

    async onGoingClick() {
        const user = await firstValueFrom(this.afAuth.authState)
        if (user == null) {
            await this.router.navigate(['sign-in'])
        }

        const partyId = this.party?.id
        if (partyId == undefined)
            return

        await this.partyService.answerParty(partyId, WatchStatus.GOING)
        this.party!.watchStatus = WatchStatus.GOING
        this.party!.isNotInterested = false
    }

    onDuplicateClick() {
        const partyId = this.party?.id
        if (partyId == undefined)
            return

        const dialogRef = this.matDialog.open(PartyDuplicateComponent, {
            panelClass: 'cheers-dialog'
        });

        dialogRef.afterClosed().subscribe(async result => {
            if (result == false)
                return;

            await firstValueFrom(this.partyService.duplicateParty(partyId));
            this.snackBar.open("Duplicated party", "Hide");
        });
    }

    onTransferClick() {
        const partyId = this.party?.id
        if (partyId == undefined)
            return

        const dialogRef = this.matDialog.open(PartyTransferComponent, {
            panelClass: 'cheers-dialog',
            data: partyId,
        });
    }

    goToGoogleMapsLink(latitude: number, longitude: number){
        const url = `https://www.google.com/maps/search/?api=1&query=${latitude},${longitude}`
        window.open(url, "_blank");
    }

    async onAnswerClick(status: WatchStatus) {
        const partyId = this.party?.id
        if (partyId == undefined)
            return

        await this.partyService.answerParty(partyId, status)
        this.party!.watchStatus = status

        if (status == WatchStatus.NOT_INTERESTED)
            this.party!.isNotInterested = true
    }
}
