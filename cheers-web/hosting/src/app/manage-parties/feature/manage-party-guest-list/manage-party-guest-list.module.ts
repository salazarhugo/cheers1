import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyGuestListRoutingModule } from './manage-party-guest-list-routing.module';
import { ManagePartyGuestListComponent } from './manage-party-guest-list.component';
import {UserItemModule} from "../../../users/ui/user-item/user-item.module";


@NgModule({
  declarations: [
    ManagePartyGuestListComponent
  ],
    imports: [
        CommonModule,
        ManagePartyGuestListRoutingModule,
        UserItemModule
    ]
})
export class ManagePartyGuestListModule { }
