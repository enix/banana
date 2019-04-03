import deepmerge from 'deepmerge';

export default function extendUpdate(update) {
  update.extend('$auto', (value, object) => {
    return object
      ? update(object, { $merge: deepmerge(object, value) })
      : update({}, { $merge: value });
  });

  update.extend('$autoArray', (value, object) => {
    const arrValue = Object.keys(value).reduce((arr, key) => {
      arr[key] = value[key];
      return arr
    }, []);

    return object
      ? update(object, { $merge: deepmerge(object, arrValue) })
      : update([], { $merge: arrValue });
  });

  return update;
};
