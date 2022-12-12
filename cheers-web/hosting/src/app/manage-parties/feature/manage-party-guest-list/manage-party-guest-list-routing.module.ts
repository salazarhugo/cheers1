import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ManagePartyGuestListComponent} from "./manage-party-guest-list.component";

const routes: Routes = [{path: '', component: ManagePartyGuestListComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartyGuestListRoutingModule { }
