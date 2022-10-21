// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Mytestlab123Iio1Iio1 from './mytestlab123.iio1.iio1'


export default { 
  Mytestlab123Iio1Iio1: load(Mytestlab123Iio1Iio1, 'mytestlab123.iio1.iio1'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}