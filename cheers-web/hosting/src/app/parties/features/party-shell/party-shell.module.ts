import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PartyShellRoutingModule } from './party-shell-routing.module';
import {PartyTransferModule} from "../party-transfer/party-transfer.module";
import {PartyDuplicateModule} from "../party-duplicate/party-duplicate.module";


@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        PartyShellRoutingModule,
        PartyTransferModule,
        PartyDuplicateModule,
    ]
})
export class PartyShellModule { }
