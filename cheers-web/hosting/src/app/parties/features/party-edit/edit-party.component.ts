import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, UntypedFormBuilder, Validators} from "@angular/forms";
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../data/party.service";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {AngularFireStorage} from "@angular/fire/compat/storage";
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {MapboxGeocodingService} from "../../../shared/data/services/mapbox-geocoding.service";
import {Party} from "../../../shared/data/models/party.model";
import {Observable, of} from "rxjs";
import {User} from "../../../shared/data/models/user.model";
import {Privacy} from "../../../shared/data/enum/privacy.enum";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
    selector: 'app-user-edit-party-detail',
    templateUrl: './edit-party.component.html',
    styleUrls: ['./edit-party.component.css', '../party-create/create-party.component.sass']
})
export class EditPartyComponent implements OnInit {

    party: Party
    partyForm = new FormGroup({
        name: new FormControl('', Validators.required),
        description: new FormControl('', {nonNullable: false}),
        startDate: new FormControl(new Date(), Validators.required),
        startTime: new FormControl(0),
        endDate: new FormControl(0),
        endTime: new FormControl(0),
        bannerUrl: new FormControl(''),
        privacy: new FormControl(Privacy.FRIENDS),
        latitude: new FormControl(0),
        longitude: new FormControl(0),
        locationName: new FormControl(''),
    });
    hasEndDate = false
    $user: Observable<User | null> = of(null)
    isLoading: boolean = false

    constructor(
        private userService: UserService,
        private partyService: PartyService,
        private _snackBar: MatSnackBar,
        private fb: UntypedFormBuilder,
        private fs: AngularFirestore,
        private storage: AngularFireStorage,
        private route: ActivatedRoute,
        private mapboxService: MapboxGeocodingService,
    ) {
        this.$user = this.userService.user$
        this.route.paramMap.subscribe((params: ParamMap) => {
            const partyId = params.get("id") as string
            this.partyService.getParty(partyId).subscribe((party: any) => {
                console.log(party)
                party.startDate = new Date(party.startDate)
                party.endDate = new Date(party.endDate)
                this.partyForm.patchValue(party)
            })
        })
    }

    ngOnInit(): void {
    }

    onSubmit() {

    }
}
