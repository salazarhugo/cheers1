import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {Party} from "../../../shared/data/models/party.model";
import {PartyService} from "../../data/party.service";
import {Clipboard} from '@angular/cdk/clipboard';
import {MatSnackBar} from "@angular/material/snack-bar";
import {PartyAnswer} from "../../../../gen/ts/cheers/party/v1/party_service";

@Component({
    selector: 'app-party-detail',
    templateUrl: './party.component.html',
    styleUrls: ['./party.component.sass']
})
export class PartyComponent implements OnInit {

    party: Party | undefined
    partyId: string | null

    constructor(
        public router: Router,
        private route: ActivatedRoute,
        private partyService: PartyService,
        private clipboard: Clipboard,
        private snackBar: MatSnackBar,
    ) {
    }

    ngOnInit(): void {
        this.route.paramMap.subscribe(async (params: ParamMap) => {
            const partyId = params.get("id")
            this.partyId = params.get("id")
            if (partyId)
                this.party = await this.partyService.getParty(partyId).toPromise()
        })
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

    async onGoingToggle() {
        const party = this.party!!
        if (party.going) {
            party.going = false
            await this.partyService.answerParty(this.partyId!, PartyAnswer.NOT_INTERESTED)
        }
        else {
            party.going = true
            await this.partyService.answerParty(this.partyId!, PartyAnswer.GOING)
        }
    }
}
