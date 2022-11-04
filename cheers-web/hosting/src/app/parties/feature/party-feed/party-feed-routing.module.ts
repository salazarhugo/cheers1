import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {PartiesComponent} from "./parties.component";

const routes: Routes = [{path: '', component: PartiesComponent}];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class PartyFeedRoutingModule {
}
