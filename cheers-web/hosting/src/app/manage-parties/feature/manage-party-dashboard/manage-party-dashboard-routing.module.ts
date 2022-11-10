import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ManagePartyDashboardComponent} from "./manage-party-dashboard.component";

const routes: Routes = [{path: '', component: ManagePartyDashboardComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartyDashboardRoutingModule { }
