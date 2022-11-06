import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ManagePartyBasicinfoComponent} from "./manage-party-basicinfo.component";

const routes: Routes = [{path: '', component: ManagePartyBasicinfoComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartyBasicinfoRoutingModule { }
