import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ManagePartyComponent} from "./manage-party.component";

const routes: Routes = [{path: '', component: ManagePartyComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartyRoutingModule { }
