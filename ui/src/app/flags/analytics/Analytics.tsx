import { useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { useOutletContext } from 'react-router-dom';
import Combobox from '~/components/forms/Combobox';
import 'chartjs-adapter-date-fns';
import { add, addMinutes, format, parseISO } from 'date-fns';
import { selectCurrentNamespace } from '~/app/namespaces/namespacesSlice';
import { IFlag } from '~/types/Flag';
import { BarGraph } from '~/components/graphs';
import { IFilterable } from '~/types/Selectable';
import { Formik } from 'formik';
import { useGetFlagEvaluationCountQuery } from '~/app/flags/analyticsApi';

type AnalyticsProps = {
  flag: IFlag;
};

const timeFormat = 'yyyy-MM-dd HH:mm:ss';

interface IDuration {
  value: number;
}

type FilterableDurations = IDuration & IFilterable;

const durations: FilterableDurations[] = [
  {
    value: 30,
    key: '30 minutes',
    displayValue: '30 minutes',
    filterValue: '30 minutes'
  },
  {
    value: 60,
    key: '1 hour',
    displayValue: '1 hour',
    filterValue: '1 hour'
  },
  {
    value: 60 * 4,
    key: '4 hours',
    displayValue: '4 hours',
    filterValue: '4 hours'
  },
  {
    value: 60 * 12,
    key: '12 hours',
    displayValue: '12 hours',
    filterValue: '12 hours'
  }
];

export default function Analytics() {
  const [selectedDuration, setSelectedDuration] =
    useState<FilterableDurations | null>(durations[0]);
  const { flag } = useOutletContext<AnalyticsProps>();
  const namespace = useSelector(selectCurrentNamespace);

  const nowISO = parseISO(new Date().toISOString());

  const getFlagEvaluationCount = useGetFlagEvaluationCountQuery({
    namespaceKey: namespace.key,
    flagKey: flag.key,
    from: format(
      addMinutes(
        addMinutes(
          nowISO,
          selectedDuration?.value ? selectedDuration.value * -1 : -60
        ),
        nowISO.getTimezoneOffset()
      ),
      timeFormat
    ),
    to: format(addMinutes(nowISO, nowISO.getTimezoneOffset()), timeFormat)
  });

  const flagEvaluationCount = useMemo(() => {
    return {
      timestamps: getFlagEvaluationCount.data?.timestamps,
      values: getFlagEvaluationCount.data?.values
    };
  }, [getFlagEvaluationCount]);

  const initialValues = {
    durationValue: selectedDuration?.key
  };

  return (
    <div className="mx-12 my-12">
      <>
        <Formik
          initialValues={initialValues}
          onSubmit={async function () {
            throw new Error('not implemented');
          }}
        >
          {() => (
            <Combobox<FilterableDurations>
              id="durationValue"
              name="durationValue"
              placeholder="Select duration"
              className="absolute right-24 z-20"
              values={durations}
              selected={selectedDuration}
              setSelected={setSelectedDuration}
            />
          )}
        </Formik>
      </>
      <div className="relative top-8">
        <BarGraph
          timestamps={flagEvaluationCount.timestamps || []}
          values={flagEvaluationCount.values || []}
          flagKey={flag.key}
        />
      </div>
    </div>
  );
}
