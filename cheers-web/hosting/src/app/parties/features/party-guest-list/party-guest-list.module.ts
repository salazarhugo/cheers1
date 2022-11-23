import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyGuestListRoutingModule} from './party-guest-list-routing.module';
import {PartyGuestListComponent} from './party-guest-list.component';


@NgModule({
    declarations: [
        PartyGuestListComponent
    ],
    imports: [
        CommonModule,
        PartyGuestListRoutingModule
    ],
    exports: [
        PartyGuestListComponent
    ],
})
export class PartyGuestListModule {
}
