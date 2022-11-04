import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ActivityItemComponent} from "./activity-item.component";

const routes: Routes = [{path: '', component: ActivityItemComponent}];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class ActivityItemRoutingModule {
}
