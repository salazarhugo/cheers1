import { Component, OnInit } from '@angular/core';
import {Party} from "../../../shared/data/models/party.model";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {Privacy} from "../../../shared/data/enum/privacy.enum";

@Component({
  selector: 'app-manage-party-basicinfo',
  templateUrl: './manage-party-basicinfo.component.html',
  styleUrls: ['./manage-party-basicinfo.component.sass']
})
export class ManagePartyBasicinfoComponent implements OnInit {

  partyForm = new FormGroup({
    name: new FormControl('', Validators.required),
    description: new FormControl(''),
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

  constructor() { }

  ngOnInit(): void {
  }

  onSubmit($event: Party) {

  }
}
