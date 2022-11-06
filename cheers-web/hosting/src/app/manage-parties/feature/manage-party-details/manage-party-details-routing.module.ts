import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ManagePartyDetailsComponent} from "./manage-party-details.component";

const routes: Routes = [{path: '', component: ManagePartyDetailsComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartyDetailsRoutingModule { }
