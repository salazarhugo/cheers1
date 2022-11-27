import {Component, OnInit} from '@angular/core';
import {Party} from "../../../shared/data/models/party.model";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {Privacy} from "../../../shared/data/enum/privacy.enum";
import {PartyService} from "../../../parties/data/party.service";
import {lastValueFrom, Observable} from "rxjs";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
    selector: 'app-manage-party-basicinfo',
    templateUrl: './manage-party-basicinfo.component.html',
    styleUrls: ['./manage-party-basicinfo.component.sass']
})
export class ManagePartyBasicinfoComponent implements OnInit {

    isLoading = false
    $party: Observable<Party>

    partyForm = new FormGroup({
        id: new FormControl(''),
        name: new FormControl('', Validators.required),
        description: new FormControl(''),
        startDate: new FormControl(new Date(), Validators.required),
        startTime: new FormControl(''),
        endDate: new FormControl(new Date()),
        endTime: new FormControl(new Date(), Validators.required),
        privacy: new FormControl(Privacy.FRIENDS),
        latitude: new FormControl(0),
        longitude: new FormControl(0),
        locationName: new FormControl(''),
    });

    constructor(
        private partyService: PartyService,
        private _snackBar: MatSnackBar,
    ) {
    }

    ngOnInit(): void {
        this.partyService.getManagedParty().subscribe(party => {
            this.partyForm.patchValue({
                id: party.id,
                name: party.name,
                description: party.description,
                startDate: new Date(party.startDate * 1000),
                startTime: new Date(party.startDate * 1000).getHours() + ":" + new Date(party.startDate * 1000).getMinutes(),
                endDate: new Date(party.endDate * 1000),
                locationName: party.locationName,
            })
        })
    }

    openSnackBar(message: string, action: string) {
        this._snackBar.open(message, action, {duration: 3000});
    }

    async onSubmit(party: Party) {
        this.isLoading = true

        console.log(party)

        try {
            const response = await lastValueFrom(this.partyService.updateParty(party))
            this.openSnackBar("Party updated", 'OK')
        } catch (e) {
            console.log(e)
            this.openSnackBar("Failed to update party", 'OK')
        }
        this.isLoading = false
    }
}
