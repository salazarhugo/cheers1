import {Component, OnInit} from '@angular/core';
import {firstValueFrom, map, Observable, of} from "rxjs";
import {Post} from "../../../shared/data/models/post.model";
import {Story} from "../../../shared/data/models/story.model";
import {UserService} from "../../../shared/data/services/user.service";
import {StoryService} from "../../../stories/data/story.service";
import {PartyService} from "../../data/party.service";
import {Party} from "../../../shared/data/models/party.model";
import {orderBy} from "@angular/fire/firestore";

@Component({
    selector: 'app-parties',
    templateUrl: './parties.component.html',
    styleUrls: ['./parties.component.sass']
})
export class PartiesComponent implements OnInit {

    $parties: Observable<Party[] | null> = of(null)
    $myParties: Observable<Party[] | null> = of(null)
    chips: string[] = ["All", "Mix", "Outside", "Indoor", "Gadgets", "Android", "Piano", "Cloud computing", "Nouveaute", "Echecs"];

    constructor(
        private userService: UserService,
        private partyService: PartyService,
    ) {
    }

    async ngOnInit() {
        const user = await firstValueFrom(this.userService.user$)
        this.$myParties = this.partyService.getMyParties()
        this.$parties = this.partyService.getPartyFeed().pipe(
            map(parties => parties.map(party => {
                party.owner = party.hostId == user.id
                console.log(party.hostId)
                return party
            })),
            map(results => results.sort((a, b) => (a.createTime < b.createTime) ? 1 : -1))
        )
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }
}
