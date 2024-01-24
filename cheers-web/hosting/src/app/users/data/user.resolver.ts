import {Injectable} from '@angular/core';
import { ActivatedRouteSnapshot, Router } from "@angular/router";
import {UserService} from '../../shared/data/services/user.service';
import {first} from "rxjs/operators";
import {UserModel} from "../../shared/data/models/user.model";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {ApiService} from "../../shared/data/services/api.service";
import {EMPTY, firstValueFrom, throwError} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class UserResolver  {

  constructor(
      public userService: UserService,
      private router: Router,
      private afAuth: AngularFireAuth,
      private api: ApiService,
  ) {
  }

  resolve(route: ActivatedRouteSnapshot): Promise<UserModel | null> {
    return new Promise(async (resolve, reject) => {

        const authUser = await firstValueFrom(this.afAuth.authState)

        if (authUser) {
            let user = await firstValueFrom(this.api.getUser(authUser.uid))

            if (user == null) {
                console.log("User doesn't exist")
                await this.router.navigate(['finish-sign-up']);
                return reject();
            }

            const token = await authUser?.getIdTokenResult()
            if (!token)
                return

            user.admin = token.claims["admin"] != null
            user.moderator = token.claims["moderator"] != null
            user.business = token.claims["business"] != null

            this.userService.setUser(user)
            return resolve(user)
        }
        else {
            return resolve(null)
        }
    })
  }
}
