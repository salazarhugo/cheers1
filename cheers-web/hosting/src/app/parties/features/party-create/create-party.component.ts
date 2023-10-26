import { Component, OnInit } from '@angular/core';
import {UntypedFormBuilder, FormControl, FormGroup, Validator, Validators} from "@angular/forms";
import {lastValueFrom} from "rxjs";
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../data/party.service";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {ActivatedRoute, Router} from "@angular/router";
import {AngularFireStorage} from "@angular/fire/compat/storage";
import {MapboxGeocodingService} from "../../../shared/data/services/mapbox-geocoding.service";
import {Privacy} from "../../../shared/data/enum/privacy.enum";
import {Party} from "../../../shared/data/models/party.model";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
    selector: 'app-create-party-detail',
    templateUrl: './create-party.component.html',
    styleUrls: ['./create-party.component.sass']
})
export class CreatePartyComponent implements OnInit {

    isLoading: boolean = false

    constructor(
        private userService: UserService,
        private partyService: PartyService,
        private _snackBar: MatSnackBar,
        private fb: UntypedFormBuilder,
        private fs: AngularFirestore,
        private storage: AngularFireStorage,
        private router: Router,
        private route: ActivatedRoute,
        private mapboxService: MapboxGeocodingService,
    ) {
    }

    partyForm = new FormGroup({
        name: new FormControl('', Validators.required),
        description: new FormControl(''),
        startDate: new FormControl(new Date(), Validators.required),
        startTime: new FormControl(''),
        endDate: new FormControl(new Date()),
        endTime: new FormControl(''),
        bannerUrl: new FormControl(''),
        privacy: new FormControl(Privacy.FRIENDS),
        latitude: new FormControl(0),
        longitude: new FormControl(0),
        locationName: new FormControl(''),
    });


    ngOnInit(): void {
    }

    openSnackBar(message: string, action: string) {
        this._snackBar.open(message, action, {duration: 3000});
    }

    async onSubmit(partyForm: any) {
        this.isLoading = true

        let party: Party = Object.assign(new Party(), partyForm);

        try {
            console.log("HUGO")
            console.log(party)
            const response = await lastValueFrom(this.partyService.createParty(party))
            this.openSnackBar("Party created", 'OK')
            await this.router.navigate(['/manage/parties', response.party?.id!, 'tickets'])
        } catch (e) {
            console.log(e)
            this.openSnackBar("Failed to create party-detail", 'OK')
        }
        this.isLoading = false
    }
}
