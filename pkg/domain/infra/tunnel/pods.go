package tunnel

import (
	"context"

	"github.com/containers/libpod/pkg/bindings/pods"
	"github.com/containers/libpod/pkg/domain/entities"
)

func (ic *ContainerEngine) PodExists(ctx context.Context, nameOrId string) (*entities.BoolReport, error) {
	exists, err := pods.Exists(ic.ClientCxt, nameOrId)
	return &entities.BoolReport{Value: exists}, err
}

func (ic *ContainerEngine) PodKill(ctx context.Context, namesOrIds []string, options entities.PodKillOptions) ([]*entities.PodKillReport, error) {
	var (
		reports []*entities.PodKillReport
	)
	foundPods, err := getPodsByContext(ic.ClientCxt, options.All, namesOrIds)
	if err != nil {
		return nil, err
	}
	for _, p := range foundPods {
		response, err := pods.Kill(ic.ClientCxt, p.Id, &options.Signal)
		if err != nil {
			report := entities.PodKillReport{
				Errs: []error{err},
				Id:   p.Id,
			}
			reports = append(reports, &report)
			continue
		}
		reports = append(reports, response)
	}
	return reports, nil
}

func (ic *ContainerEngine) PodPause(ctx context.Context, namesOrIds []string, options entities.PodPauseOptions) ([]*entities.PodPauseReport, error) {
	var (
		reports []*entities.PodPauseReport
	)
	foundPods, err := getPodsByContext(ic.ClientCxt, options.All, namesOrIds)
	if err != nil {
		return nil, err
	}
	for _, p := range foundPods {
		response, err := pods.Pause(ic.ClientCxt, p.Id)
		if err != nil {
			report := entities.PodPauseReport{
				Errs: []error{err},
				Id:   p.Id,
			}
			reports = append(reports, &report)
			continue
		}
		reports = append(reports, response)
	}
	return reports, nil
}

func (ic *ContainerEngine) PodUnpause(ctx context.Context, namesOrIds []string, options entities.PodunpauseOptions) ([]*entities.PodUnpauseReport, error) {
	var (
		reports []*entities.PodUnpauseReport
	)
	foundPods, err := getPodsByContext(ic.ClientCxt, options.All, namesOrIds)
	if err != nil {
		return nil, err
	}
	for _, p := range foundPods {
		response, err := pods.Unpause(ic.ClientCxt, p.Id)
		if err != nil {
			report := entities.PodUnpauseReport{
				Errs: []error{err},
				Id:   p.Id,
			}
			reports = append(reports, &report)
			continue
		}
		reports = append(reports, response)
	}
	return reports, nil
}

func (ic *ContainerEngine) PodStop(ctx context.Context, namesOrIds []string, options entities.PodStopOptions) ([]*entities.PodStopReport, error) {
	var (
		reports []*entities.PodStopReport
		timeout int = -1
	)
	foundPods, err := getPodsByContext(ic.ClientCxt, options.All, namesOrIds)
	if err != nil {
		return nil, err
	}
	if options.Timeout != -1 {
		timeout = options.Timeout
	}
	for _, p := range foundPods {
		response, err := pods.Stop(ic.ClientCxt, p.Id, &timeout)
		if err != nil {
			report := entities.PodStopReport{
				Errs: []error{err},
				Id:   p.Id,
			}
			reports = append(reports, &report)
			continue
		}
		reports = append(reports, response)
	}
	return reports, nil
}

func (ic *ContainerEngine) PodRestart(ctx context.Context, namesOrIds []string, options entities.PodRestartOptions) ([]*entities.PodRestartReport, error) {
	var reports []*entities.PodRestartReport
	foundPods, err := getPodsByContext(ic.ClientCxt, options.All, namesOrIds)
	if err != nil {
		return nil, err
	}
	for _, p := range foundPods {
		response, err := pods.Restart(ic.ClientCxt, p.Id)
		if err != nil {
			report := entities.PodRestartReport{
				Errs: []error{err},
				Id:   p.Id,
			}
			reports = append(reports, &report)
			continue
		}
		reports = append(reports, response)
	}
	return reports, nil
}

func (ic *ContainerEngine) PodStart(ctx context.Context, namesOrIds []string, options entities.PodStartOptions) ([]*entities.PodStartReport, error) {
	var reports []*entities.PodStartReport
	foundPods, err := getPodsByContext(ic.ClientCxt, options.All, namesOrIds)
	if err != nil {
		return nil, err
	}
	for _, p := range foundPods {
		response, err := pods.Start(ic.ClientCxt, p.Id)
		if err != nil {
			report := entities.PodStartReport{
				Errs: []error{err},
				Id:   p.Id,
			}
			reports = append(reports, &report)
			continue
		}
		reports = append(reports, response)
	}
	return reports, nil
}

func (ic *ContainerEngine) PodRm(ctx context.Context, namesOrIds []string, options entities.PodRmOptions) ([]*entities.PodRmReport, error) {
	var reports []*entities.PodRmReport
	foundPods, err := getPodsByContext(ic.ClientCxt, options.All, namesOrIds)
	if err != nil {
		return nil, err
	}
	for _, p := range foundPods {
		response, err := pods.Remove(ic.ClientCxt, p.Id, &options.Force)
		if err != nil {
			report := entities.PodRmReport{
				Err: err,
				Id:  p.Id,
			}
			reports = append(reports, &report)
			continue
		}
		reports = append(reports, response)
	}
	return reports, nil
}