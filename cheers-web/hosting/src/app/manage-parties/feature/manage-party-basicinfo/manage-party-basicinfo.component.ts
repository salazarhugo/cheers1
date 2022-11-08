import {Component, OnInit} from '@angular/core';
import {Party} from "../../../shared/data/models/party.model";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {Privacy} from "../../../shared/data/enum/privacy.enum";
import {PartyService} from "../../../parties/data/party.service";
import {Observable} from "rxjs";

@Component({
    selector: 'app-manage-party-basicinfo',
    templateUrl: './manage-party-basicinfo.component.html',
    styleUrls: ['./manage-party-basicinfo.component.sass']
})
export class ManagePartyBasicinfoComponent implements OnInit {

    $party: Observable<Party>

    partyForm = new FormGroup({
        name: new FormControl('', Validators.required),
        description: new FormControl(''),
        startDate: new FormControl(new Date(), Validators.required),
        startTime: new FormControl(0),
        endDate: new FormControl(0),
        endTime: new FormControl(0),
        latitude: new FormControl(0),
        longitude: new FormControl(0),
        locationName: new FormControl(''),
    });

    constructor(
        private partyService: PartyService,
    ) {
    }

    ngOnInit(): void {
        this.partyService.getManagedParty().subscribe(party => {
            this.partyForm.patchValue({
                name: party.name,
                description: party.description,
                startDate: new Date(party.startDate),
                locationName: party.locationName,
            })
        })
    }

    onSubmit($event: Party) {

    }
}
