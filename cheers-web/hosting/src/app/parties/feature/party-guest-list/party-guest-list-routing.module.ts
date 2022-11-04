import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PartyGuestListComponent} from "./party-guest-list.component";

const routes: Routes = [{path: '', component: PartyGuestListComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PartyGuestListRoutingModule { }
