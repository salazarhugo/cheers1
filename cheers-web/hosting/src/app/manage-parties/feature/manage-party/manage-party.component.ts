import { Component, OnInit } from '@angular/core';
import {fadeAnimation} from "../../../animtaions";

@Component({
  selector: 'app-manage-party',
  templateUrl: './manage-party.component.html',
  styleUrls: ['./manage-party.component.sass'],
  animations: [fadeAnimation],
})
export class ManagePartyComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
