import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyGuestListRoutingModule } from './manage-party-guest-list-routing.module';
import { ManagePartyGuestListComponent } from './manage-party-guest-list.component';


@NgModule({
  declarations: [
    ManagePartyGuestListComponent
  ],
  imports: [
    CommonModule,
    ManagePartyGuestListRoutingModule
  ]
})
export class ManagePartyGuestListModule { }
