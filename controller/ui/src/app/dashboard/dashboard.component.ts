import { FormBuilder, FormControl, FormArray } from '@angular/forms'
import { Component, OnInit } from '@angular/core'

import { Proxy } from '../models/proxy'

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  dashboardForm = this._fb.group({
    proxies: this._fb.array([])
  })

  constructor(private _fb: FormBuilder) {}

  ngOnInit() {
  	this.addProxy(new Proxy)
  }

  get proxies() {
    return this.dashboardForm.get('proxies') as FormArray;
  }

  addProxy(proxy: Proxy) {
    console.log(this.proxies);
    this.proxies.push(this._fb.group({
      port: new FormControl(proxy.port),
      env_base_path: new FormControl(proxy.env_base_path),
      site_base_path: new FormControl(proxy.site_base_path)
    }));
  }
}
